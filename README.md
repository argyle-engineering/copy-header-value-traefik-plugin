# Copy header value Traefik plugin

Traefik plugin that copies HTTP header value with format `key1=value1; key2=value2` into a new 
header. Motivation for this plugin is to be able to extract a particular key's value from Cookie 
header. However, it can be adapted for other purposes taking into account this plugin's 
limitations, e.g. only a single key's value can be copied.

### Configuration

Traefik static configuration for local plugin:

```.yaml
...
experimental:
  plugins:
    copy-header-value:
      moduleName: github.com/argyle-engineering/copy-header-value-traefik-plugin
```

Plugin is then configured as a route middleware

```.yaml
http:
  routers:
    route1:
      middlewares: ["copyHeaderValue"]
  middlewares:
    copyHeaderValue:
      plugin:
        copy-header-value:
          from: 'Cookie'
          pairSeparator: ';'            
          keyValueSeparator: '='
          key: 'id'
          to: 'Authorization'
          prefix: 'Bearer '
          overwrite: false
```

In this example, when there is `Cookie` header, value identified by the `id` key will be moved
into `Authorization` header with `Bearer ` as a value prefix. In case `Authorization` header 
already exists, `overwrite` controls whether it will be overwritten with the value extracted from 
`Cookie` header.
