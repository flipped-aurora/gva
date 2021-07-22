// +build mysql

package config

func (g *Gorm) GetDsn() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].Username + ":" + g.Dsn.Sources[0].Password + "@tcp(" + g.Dsn.Sources[0].Host + ":" + g.Dsn.Sources[0].Port + ")/" + g.Dsn.Sources[0].DbName + "?" + g.Config
	}
	return ""
}

func (g *Gorm) GetDbName() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].DbName
	}
	return ""
}

func (g *Gorm) GetDatabaseDsn() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].Username + ":" + g.Dsn.Sources[0].Password + "@tcp(" + g.Dsn.Sources[0].Host + ":" + g.Dsn.Sources[0].Port + ")/"
	}
	return ""
}
