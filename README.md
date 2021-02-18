# go-server-crud
 tutorial backend rest api work with frontend vue.js

<table border=none>
  <tr align=left>
    <td>Framework By  <a href="https://github.com/gin-gonic/gin">Gin Web Framework</a></td>
    <td><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="10%"></td>
</tr>
</table>







method | url | request | response
--- | --- | --- | ---
GET | /api/tutorials | | [{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]
GET | /api/tutorials/602aa1e04f3b51804eca6917 ||{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}
POST | /api/tutorials | {"title":"xx","description":"xx Description"} | Inserted a single document Success
PUT | /api/tutorials | {"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"} | Updated  a single document Success
DELETE | /api/tutorials/602aa1e04f3b51804eca6917 ||Deleted id:602aa1e04f3b51804eca6917
DELETE | /api/tutorials ||All deleted
