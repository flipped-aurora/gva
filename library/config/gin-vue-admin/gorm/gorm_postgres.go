// +build postgres

package config

func (g *Gorm) GetDsn() string {
	if len(g.Dsn.Sources) >= 1 {
		return "host=" + g.Dsn.Sources[0].Host + " user=" + g.Dsn.Sources[0].Username + " password=" + g.Dsn.Sources[0].Password + " dbname=" + g.Dsn.Sources[0].DbName + " port=" + g.Dsn.Sources[0].Port + " " + g.Config
	}
	return ""
}

func (g *Gorm) GetDatabaseDsn() string {
	if len(g.Dsn.Sources) >= 1 {
		return "host=" + g.Dsn.Sources[0].Host + " user=" + g.Dsn.Sources[0].Username + " password=" + g.Dsn.Sources[0].Password + " sslmode=disable"
	}
	return ""
}

func (g *Gorm) GetDbName() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].DbName
	}
	return ""
}

func (g *Gorm) GetHost() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].Host
	}
	return ""
}

func (g *Gorm) GetUsername() string {
	if len(g.Dsn.Sources) >= 1 {
		return g.Dsn.Sources[0].Username
	}
	return ""
}
