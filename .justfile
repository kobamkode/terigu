start:
    rm -rf ./templates
    npm run build
    mkdir -p ./templates/assets/
    mv ./templates/*.css ./templates/assets/
    mv ./templates/*.map ./templates/assets/
    RUST_LOG=debug cargo run
