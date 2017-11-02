# nebula-dal-generator
> nebula data access layer code generator

## 使用最讨巧最简单的办法
> 配置文件的参数使用`:field_name`，举例如下:
    
```
    配置文件的SQL语句：
    SELECT id, app_id, api_hash, title, short_name FROM apps WHERE id=:id
    
    通过:id解析出参数，自动生成入参代码
```
