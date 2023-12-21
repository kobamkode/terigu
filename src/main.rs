use askama::Template;
use axum::{Router, routing::get};
use tokio::net::TcpListener;
use tower_http::services::ServeDir;

#[derive(Template)]
#[template(path = "index.html")]
struct HelloTemplate<'a> {
    name: &'a str,
}

async fn hello() -> HelloTemplate<'static> {
    HelloTemplate { name: "world" }
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    let app = Router::new().nest_service("/assets", ServeDir::new("templates/assets"))
        .route("/", get(hello));

    let listener = TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
