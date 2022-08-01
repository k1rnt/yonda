# book seed
1. `curl -XPOST -H "Content-type: application/json" -d '{"name":"実践ドメイン駆動設計", "author":"Vaughn Vernon" ,"desc":"エリック・エヴァンスが確立した理論を実際 の設計に応用する", "page_count":583}' 'http://localhost:3000/api/v1/book/register'`
2. `curl -XPOST -H "Content-type: application/json" -d '{"name":"ハッカーと画家", "author":"Paul Graham" ,"desc":"コンピュータ時代の創造者たち", "page_count":266}' 'http://localhost:3000/api/v1/book/register'`
