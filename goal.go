package goal

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/glebarez/sqlite"
	"github.com/huoyijie/goal/auth"
	"github.com/huoyijie/goal/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var models = []any{
	&auth.User{},
	&auth.Role{},
	&auth.Session{},
}

type Item struct {
	Name string
	CanAdd,
	CanDelete,
	CanChange,
	CanGet bool
}

type Group struct {
	Name  string
	Items []*Item
}

func groupList() (groups []*Group) {
	dict := map[string][]*Item{}

	for _, model := range models {
		elem := reflect.TypeOf(model).Elem()
		pkgName := cases.Title(language.Und).String(filepath.Base(elem.PkgPath()))
		if items, ok := dict[pkgName]; ok {
			dict[pkgName] = append(items, &Item{
				Name: elem.Name(),
			})
		} else {
			dict[pkgName] = []*Item{{Name: elem.Name()}}
		}
	}

	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		items := dict[v]
		sort.Slice(items, func(i, j int) (less bool) {
			less = items[i].Name < items[j].Name
			return
		})
		groups = append(groups, &Group{v, items})
	}
	return
}

func Register(modelList ...any) {
	models = append(models, modelList)
}

func Models() []any {
	return models
}

func HomeDir() (homeDir string) {
	homeDir, err := os.UserHomeDir()
	util.LogFatal(err)
	return
}

func WorkDir() (workDir string) {
	workDir = filepath.Join(HomeDir(), ".goal")
	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		util.LogFatal(os.Mkdir(workDir, 00744))
	}
	return
}

var db *gorm.DB
var enforcer *casbin.Enforcer

//go:embed templates/*
var tmplFS embed.FS

//go:embed rbac_model.conf
var rbacModel string

func newTemplate() (tmpl *template.Template) {
	tmpl = template.Must(template.New("").ParseFS(tmplFS, "templates/*.htm"))
	return
}

func clearSessions(db *gorm.DB) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		util.LogFatal(db.Delete(&auth.Session{}, "expire_date < ?", time.Now()).Error)
	}
}

func OpenDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(filepath.Join(WorkDir(), "db.sqlite3")), &gorm.Config{})
	util.LogFatal(err)

	util.LogFatal(db.AutoMigrate(Models()...))

	go clearSessions(db)
	return
}

func Run(port int, host string) {
	model, err := model.NewModelFromString(rbacModel)
	util.LogFatal(err)

	db = OpenDB()
	adapter, err := gormadapter.NewAdapterByDB(db)
	util.LogFatal(err)

	enforcer, err = casbin.NewEnforcer(model, adapter)
	util.LogFatal(err)
	util.LogFatal(enforcer.LoadPolicy())

	router := newRouter()
	router.Run(fmt.Sprintf("%s:%d", host, port))
}
