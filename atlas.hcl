data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/migrate/main.go",
    "migrate",
    "--load",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/14/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}