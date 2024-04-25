package config

#DBConfig: {
    user: string
    password: string
    host: string
    port: int
}

#MetricConfig: {
    host: string
    port: int
}

db: #DBConfig & {
    user: "db_user"
    password: "password"
    host: "127.0.0.1"
    port: 3306
}

metric: #MetricConfig & {
    host: "http://localhost"
    port: 9091
}
