# 数据库技术基础与应用开发 - 项目增强文档

为了满足结课论文中对 MySQL 和 GORM 掌握程度的要求，本项目在原有基础上进行了以下增强：

## 1. 数据库关系建模 (E-R Modeling)
- **修改文件**: `models/user.go`, `models/article.go`
- **说明**: 建立了 `User` (用户) 和 `Article` (文章) 之间的一对多 (1:N) 关系。
    - 在 `Article` 模型中添加了 `UserID` 外键和 `User` 关联字段。
    - 在 `User` 模型中添加了 `Articles` 字段。
- **意义**: 体现了对关系型数据库核心概念（外键、关联）的理解。

## 2. 数据生命周期管理 (Hooks)
- **修改文件**: `models/user.go`, `controllers/auth_controller.go`
- **说明**: 使用 GORM 的 `BeforeSave` 钩子函数自动处理密码加密。
    - 移除了 Controller 层的手动加密逻辑，将其下沉到 Model 层。
- **意义**: 增强了数据的安全性和一致性，体现了对 ORM 框架高级特性（Hooks/Triggers）的应用。

## 3. 事务处理 (Transactions)
- **修改文件**: `controllers/article_controller.go`
- **说明**: 在创建文章 (`CreateArticle`) 时引入了数据库事务。
    - 使用 `tx := global.Db.Begin()` 开启事务，`tx.Commit()` 提交，`tx.Rollback()` 回滚。
- **意义**: 保证了操作的原子性 (Atomicity)，符合 ACID 原则，是数据库开发中的关键技术点。

## 4. 索引优化 (Indexing)
- **修改文件**: `models/exchange_rate.go`
- **说明**: 为 `ExchangeRate` 表的 `FromCurrency` 和 `ToCurrency` 字段添加了联合索引 (`idx_currency_pair`)。
- **意义**: 提高了按货币对查询汇率的效率，体现了对数据库性能优化的考量。

