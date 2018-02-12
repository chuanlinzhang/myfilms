package films

import "gopkg.in/mgo.v2"

/*
影片信息
 */
type Films struct {
	FilmsNo    string  `bson:"films_no"`//影片编号
	NameCh     string  `bson:"name_ch"`//中文片名
	NameEn     string  `bson:"name_en"`//英文片名
	Director   string  `bson:"director"`//导演
	Actor      string  `bson:"actor"`//主演
	Intro      string  `bson:"intro"`//影片简介
	FilmLength int32   `bson:"film_length"`//片长
	Pricefull  float64 `bson:"pricefull"`//全票价
	Pricest    float64 `bson:"pricest"`//学生票价
	Available  bool  `bson:"available"`//该影片资料是否有效
	Latest     string  `bson:"latest"`//最后一次更新时间
}
var collectionFilms *mgo.Collection
func init() {
	collectionFilms=DB.C("films")
}