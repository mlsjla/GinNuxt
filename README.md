# GinNuxt
## DEMO [https://www.ginnuxt.com/](https://www.ginnuxt.com/) 

- support SSR and Single Page
- use Gin and Nuxt3

# Server

> RBAC scaffolding based on GIN + Gorm 2.0 + CASBIN + WIRE (DI).

- Server Base on [gin-admin](https://github.com/LyricTian/gin-admin) 

## Features

- Follow the `RESTful API` design specification
- Use `Casbin` to implement fine-grained access to the interface design
- Use `Wire` to resolve dependencies between modules
- Provides rich `Gin` middlewares (JWTAuth,CORS,RequestLogger,RequestRateLimiter,TraceID,CasbinEnforce,Recover,GZIP)
- Support `Swagger`

## Dependent Tools

```bash
go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag
```

- [air](https://github.com/cosmtrek/air) -- Live reload for Go apps
- [wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go
- [swag](https://github.com/swaggo/swag) -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## Dependent Library

- [Gin](https://gin-gonic.com/) -- The fastest full-featured web framework for Go.
- [GORM](https://gorm.io/) -- The fantastic ORM library for Golang
- [Casbin](https://casbin.org/) -- An authorization library that supports access control models like ACL, RBAC, ABAC in Golang
- [Wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go

## Getting Started

```bash
git clone https://github.com/mlsjla/GinNuxt

cd server

go run cmd/main/main.go web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml

# Or use Makefile: make start
```

> The database and table structure will be automatically created during the startup process. After the startup is successful, you can access the swagger address through the browser: [http://127.0.0.1:10088/swagger/index.html](http://127.0.0.1:10088/swagger/index.html)

### Generate `swagger` documentation

```bash
swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/app/swagger

# Or use Makefile: make swagger
```

### Use `wire` to generate dependency injection

```bash
wire gen ./internal/app

# Or use Makefile: make wire
```

## Project Layout

```text
├── cmd
│   └── gin-nuxt-server
│       └── main.go       
├── configs
│   ├── config.toml       
│   ├── menu.yaml         
│   └── model.conf        
├── docs                  
├── internal
│   └── app
│       ├── api           
│       ├── config        
│       ├── contextx      
│       ├── dao           
│       ├── ginx          
│       ├── middleware    
│       ├── module        
│       ├── router        
│       ├── schema        
│       ├── service       
│       ├── swagger       
│       ├── test          
├── pkg
│   ├── auth              
│   │   └── jwtauth       
│   ├── errors            
│   ├── gormx             
│   ├── logger            
│   │   ├── hook
│   └── util              
│       ├── conv         
│       ├── hash         
│       ├── json
│       ├── snowflake
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts               
```

# Web
Web based on - [Nuxt3](https://v3.nuxtjs.org/docs/deployment).
We recommend to look at the [documentation](https://v3.nuxtjs.org).

## Setup

Make sure to install the dependencies

```bash
yarn install
```

## Development

Start the development server on http://localhost:3000

```bash
yarn dev
```

## Production

Build the application for production:

```bash
yarn build
```

Checkout the [deployment documentation](https://v3.nuxtjs.org/docs/deployment).


# Thanks

- [gin-admin](https://github.com/LyricTian/gin-admin).
- [deployment documentation](https://v3.nuxtjs.org/docs/deployment).
- [vditor](https://github.com/Vanessa219/vditor).