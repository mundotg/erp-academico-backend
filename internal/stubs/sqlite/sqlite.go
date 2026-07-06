package sqlite

type Dialector struct{ DSN string }

func Open(dsn string) Dialector { return Dialector{DSN: dsn} }
