package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hithub.com/JIeeiroSst/go-app/domain"
	"hithub.com/JIeeiroSst/go-app/log"
)

var (
	mutex    sync.Mutex
	instance *MysqlConn
)

type MysqlConn struct {
	db *gorm.DB
}

type Config struct {
	DSN string
}

func GetMysqlConnInstance(c *Config) *MysqlConn {
	logger := log.InitZapLog()
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := c.DSN
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				logger.Sugar().Errorf("running server mysql serror %s", err)
			}
			logger.Info("server running get mysql success")
			instance = &MysqlConn{db: db}
			err = db.AutoMigrate(&domain.Article{}, &domain.Author{})
			if err != nil {
				logger.Error("error connect db")
			}
			logger.Info("conecting db success")
		}
	}
	return instance
}

func NewMysqlConn(c *Config) *MysqlConn {
	return &MysqlConn{
		db: GetMysqlConnInstance(c).db,
	}
}

func (mysql *MysqlConn) CreateArticle(article domain.Article) (err error) {
	err = mysql.db.Create(&article).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("create err %s", err)
	}
	log.InitZapLog().Info("create sucesss")
	return
}

func (mysql *MysqlConn) ArticleAll() (articles []domain.Article, err error) {
	err = mysql.db.Find(&articles).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("find all error %s", err)
	}
	log.InitZapLog().Info("find all sucess")
	return
}

func (mysql *MysqlConn) ArticleById(id int) (article domain.Article, err error) {
	err = mysql.db.Preload("Author").Where("id = ?", id).First(&article).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("find by id %d:%s error", id, err)
	}
	log.InitZapLog().Sugar().Infof("find by id %d sucess", id)
	return
}

func (mysql *MysqlConn) UpdateArticle(id int, article domain.Article) (err error) {
	err = mysql.db.Model(domain.Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("update error %s", err)
	}
	log.InitZapLog().Sugar().Infof("update sucess %d", id)
	return
}

func (mysql *MysqlConn) DeleteArticle(id int) (err error) {
	err = mysql.db.Where("id = ?", id).Delete(&domain.Article{}).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("delete error %d:%s", id, err)
	}
	log.InitZapLog().Sugar().Infof("delete sucess %d", id)
	return
}

func (mysql *MysqlConn) CreateAuthor(author domain.Author) (err error) {
	err = mysql.db.Create(&author).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("create err %s", err)
	}
	log.InitZapLog().Info("create sucesss")
	return
}

func (mysql *MysqlConn) AuthorAll() (authors []domain.Author, err error) {
	err = mysql.db.Find(&authors).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("find all error %s", err)
	}
	log.InitZapLog().Info("find all sucess")
	return
}

func (mysql *MysqlConn) AuthorById(id int) (author domain.Author, err error) {
	err = mysql.db.Where("id = ?", id).Find(&author).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("find by id %d:%s error", id, err)
	}
	log.InitZapLog().Sugar().Infof("find by id %d sucess", id)
	return
}

func (mysql *MysqlConn) UpdateAuthor(id int, author domain.Author) (err error) {
	err = mysql.db.Model(domain.Article{}).Where("id = ?", id).Updates(author).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("update error %s", err)
	}
	log.InitZapLog().Sugar().Infof("update sucess %d", id)
	return
}

func (mysql *MysqlConn) DeleteAuthor(id int) (err error) {
	err = mysql.db.Where("id = ?", id).Delete(&domain.Author{}).Error
	if err != nil {
		log.InitZapLog().Sugar().Errorf("delete error %d:%s", id, err)
	}
	log.InitZapLog().Sugar().Infof("delete sucess %d", id)
	return
}
