## twitter-clone

### users
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| id | int(11) | No | primary |  |  |  |
| name | varchar(32) | No |  |  |  |  |
| password | char(128) | No |  |  |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| updated_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) | No |  |  |  |  |

### messages
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| deleted_at | datetime(6) | No |  |  |  |  |
| id | varchar(36) | No | primary |  |  |  |
| user_id | int(11) | No |  |  |  |  |
| content | text | No |  |  |  |  |
| is_pinned | boolean | No |  | false |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| updated_at | datetime(6) | No |  |  |  |  |

### follow
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| target_user_id | int(11) | No |  |  |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) | No |  |  |  |  |
| id | int(11) | No | primary |  |  |  |
| user_id | int(11) | No |  |  |  |  |

### favorite
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) | No |  |  |  |  |
| id | int(11) | No | primary |  |  |  |
| user_id | int(11) | No |  |  |  |  |
| target_message_id | varchar(36) | No |  |  |  |  |
