# Example Go Project

This project is an example for learning Golang design pattern and structuring folder by using OOP knowledge under the hood

## Quick Start
just type this command in CLI: `make run`

## TODO
  - <s>connect to database</s>
  - <s>use Gorm's hooks (Auditable)</s>
  - <s>prevent Gorm to keep continuously adding conditions to it by using Session in repository struct</s>
  - <s>graceful shutdown</s>
  - <s>custom recovery middleware</s>
  - <s>response writer middleware</s>
  - <s>handle errors by both gin and internal system convention</s>
  - [returning data from modified rows](https://gorm.io/docs/update.html#Returning-Data-From-Modified-Rows)
  - custom model binding and validation [(*must bind* case)](https://chenyitian.gitbooks.io/gin-web-framework/content/docs/17.html)
  - custom validation
  - add example code for data validation
  - log request body
  - logging to Sentry (optional)
  - auth middleware
  - implement `RequestContext`
  - request timeout middleware
  - auto route mapping *(SUPER HARD)
  - refactor response util (use polymorphism)
  - ...

## Related Article
https://medium.com/@anewnurse/go-series-2-create-restapi-using-gin-and-gorm-9c24f3e37324 \
https://medium.com/@anewnurse/go-series-3-domain-driven-design-ddd-f12707c68ea6
