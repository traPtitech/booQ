## DB schema

### items
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |
 | name | varchar(20) | NO |  |  |  |  |
 | code | int |  |  |  |  | ISBNコードとか物品管理コードとか |
 | owner_id | int | NO |  |  |  | 所有者のユーザーID |
 | description | text |  |  |  |  | 物品の説明文 |
 | img_url | text |  |  |  |  | 画像のURL(外部に頼る) |
 | del_flg | boolean | NO |  |  |  | 削除されてないかどうか |
 | created_at | datetime | NO |  |  |  |  |
 | updated_at | datetime | NO |  |  |  |  |

### comments
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | int | NO |  |  |  |  |
 | comment | text | NO |  |  |  |  |
 | del_flg | boolean | NO |  |  |  |  |
 | created_at | datetime | NO |  |  |  |  |
 | updated_at | datetime |  |  |  |  |  |

### users
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | name | varchar(20) | NO | UNI |  |  | 同じ名前はありえない(はず) |
 | authority | boolean | NO |  |  |  | 特権ユーザー的なやつ |

### histories
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | item_id | int | NO |  |  |  |  |
 | user_id | varchar(20) | NO |  |  |  |  |
 | type | int | NO |  |  |  | 0:返した(フリー) 1:予約(借りたい) 2:貸した(貸し出し中) 初期値は0 |
 | created_at | datetime | NO |  |  |  |  |
 | deleted_at | int |  |  |  |  | 消えるというよりは取り消したかどうか、こうすると取り消した日時も格納できるね |

### tags
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | name | varchar(20) | NO | UNI |  |  |  |

### tagmaps
 | Name | Type | NULL | Key | Default | Extra | 説明 |
 | --- | --- | --- | --- | --- | --- | --- |
 | id | int | NO | PRI |  |  |  |
 | tag_id | int | NO |  |  |  |  |
 | item_id | int | NO |  |  |  |

