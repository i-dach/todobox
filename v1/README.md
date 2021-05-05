======================
# todo について
======================

## ＜最低限実装しなければいけないもの＞
▼ RESTfulなAPIであること  
▼ Todo の要素として次の項目を保持すること  
　- タイトル  
　- 内容  
▼ 検索、登録、更新、削除の操作が行える  
▼ テストコードが書かれている  
▼ API 仕様書が書かれている  

## ＜構成について＞

https://docs.google.com/presentation/d/1NUe1vfAnCuT-Eu9VJaYqD36-1Sx0GsJPqD2S8CYwnHo/edit?usp=sharing

## ＜作業計画＞

- [x] とりあえず動くURIを作成
- [x] 裏側の挙動を作成
- [x] テストコードの作成 (本来は順番が逆だが)
- [x] TDDにてリファクタリング
- [] openAPIでserver部分を代用
- [] UIも作ってみてユニットテストをしてみる

======================
# API
======================

## ＜URI一覧＞

以下にあるURIはすべて8080ポートにてやりとりを行う。

```
GET    /helth                    --> ヘルスチェック用のURI
POST   /todo/task                 --> TODOにあるタスクを追加するようのURI
PATCH  /todo/task                 --> TODOにあるタスクの内容を変更する際に利用するURI
GET    /todo/task                 --> 完了していないTODOを取得するためのURI
POST   /todo/task/done/:id        --> TODOにあるtaskをdoneするためのURI
DELETE /todo/task/delete/:id      --> TODOにあるtaskをdeleteするためのURI
```

## ＜Reference＞
### □ [GET] /helth

ヘルスチェック用のURI

#### > Parameter

No Params

#### > Response

* Code  
    * 200
* Description
    * ヘルスチェックを行う
* E.G

    ```
    {
        "message":"helth check ok"
    }
    ```

### □ [POST]   /todo/task

TODOにタスクを追加する用のURI

#### > Parameter

No Params

#### > Request body

e.g.

```
{
    "title": "Task title...",
}
```

#### > Response

* Code  
    * 200
* Description
    * タスクの登録を行う
* E.G

    ```
    {
        "todo":"insert recode :taskid"
    }
    ```

### □ [PATCH]  /todo/task 
タスク情報更新用URI

#### > Parameter

No Params

#### > Request body

e.g.

```
{
    "id": "taskid",
    "title": "change task name",
    "description": "task description",
}
```

#### > Response

* Code  
    * 200
* Description
    * 指定IDのタスク情報を変更する
* E.G

    ```
    {
        "todo":"update recode: task id"
    }
    ```

### □ [GET]    /todo/task                 
完了していないTODOを取得するためのURI
#### > Parameter

No Params

#### > Response

* Code  
    * 200
* Description
    * 有効なタスク一覧を取得する
* E.G

    ```
    {
        "todo": {
            {
                "id": "xxx",
                "title": "titile",
                "description": "xxxx",
            },
            {
                "id": "xxx",
                "title": "titile",
                "description": "xxxx",
            },
        }
    }
    ```

### □ [POST]   /todo/task/done/:id 
TODOにあるtaskをdoneするためのURI
#### > Parameter

:id -> taskのID

#### > Request body

No Body

#### > Response

* Code  
    * 200
* Description
    * 指定IDのタスクを完了状態にする
* E.G

    ```
    {
        "todo":"done task :task id"
    }
    ```

### □ [DELETE] /todo/task/delete/:id 
TODOにあるtaskをdeleteするためのURI

#### > Parameter

:id -> taskのID

#### > Response

* Code  
    * 200
* Description
    * xxx
* E.G

    ```
    {
        "todo":"delete recode :task id"
    }
    ```
