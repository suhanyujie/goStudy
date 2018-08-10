{{define "T.navbar.tpl"}}
<nav class="navbar navbar-default">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="#">Brand</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li {{if .IsIndex}}class="active"{{end}}><a href="/">首页</a></li>
        <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
        <li><a href="https://beego.me/docs/" target="_blank">beego文档</a></li>
        <li class="dropdown" {{if .IsCategory}}class="active"{{end}}>
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">分类 <span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="https://beego.me/docs/">beego文档</a></li>
            <li><a href="#">Another action</a></li>
            <li><a href="#">Something else here</a></li>
          </ul>
        </li>
      </ul>
      <form class="navbar-form navbar-left">
        <div class="form-group">
          <input type="text" class="form-control" placeholder="Search">
        </div>
        <button type="submit" class="btn btn-default">Submit</button>
      </form>
      <ul class="nav navbar-nav navbar-right">
      {{if .HasLogin}}
        <li><a href="/logout">登出</a></li>
      {{else}}
        <li><a href="/login">登录</a></li>
      {{end}}

      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>

{{end}}