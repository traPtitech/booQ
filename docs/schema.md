## DB schema

### items
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | **id** | int | NO | PRI |  |
 | name | varchar(64) | NO |  |  |  | 物品名 |
 | type | int | NO |  |  |  | 物品のタイプ(0: 備品, 1: 本(備品以外)) |
 | code | text |  |  |  |  | ISBNコードとか物品管理コードとか |
 | description | text |  |  |  |  | 物品の説明文 |
 | img_url | text |  |  |  |  | 画像のURL(外部に頼る) |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### comments
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | **id** | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | int | NO |  |  |  |  |
 | text | text | NO |  |  |  | コメントの中身 |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### users
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | **id** | int | NO | PRI |  |  |  |
 | name | varchar(32) | NO | UNI |  |  | 同じ名前はありえない(はず) |
 | displayName | varchar(64) | NO |  |  |  |  |
 | iconFileID | uuid(varchar(36)) | NO |  |  |  |  |
 | admin | boolean | NO |  | false |  | 特権ユーザー的なやつ |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |
 

### logs
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | **id** | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | int | NO |  |  |  | アクションを起こす人 |
 | owner_id | int | NO |  |  |  | 物品の所有者 |
 | type | int | NO |  |  |  | 0:借りた, 1:返した  |
 | purpose | text |  |  |  |  | 借りる目的 |
 | due_date | datetime |  |  |  |  | 返却予定日 |
 | **created_at** | datetime | NO |  |  |  |  |
 | **update_at** | datetime | NO |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### tags
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | **id** | int | NO | PRI |  |  |  |
 | name | varchar(32) | NO | UNI |  |  |  |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### tagmaps
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | tag_id | int | NO |  |  |  |  |
 | item_id | int | NO |  |  |  |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### ownershipmaps
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | int | NO |  |  |  |
 | rentalable | boolean | NO |  | true |  | 今貸し出し可能か否か |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |

### likemaps
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | int | NO |  |  |  |
 | **created_at** | datetime | NO |  |  |  |  |
 | **updated_at** | datetime |  |  |  |  |  |
 | **deleted_at** | datetime |  |  |  |  |  |