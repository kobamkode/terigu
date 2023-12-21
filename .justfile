start:
    rm -rf templates
    pnpm run build
    mkdir -p templates/assets/
    mv templates/*.css templates/assets/
    mv templates/*.map templates/assets/
    RUST_LOG=debug cargo run
