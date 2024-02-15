{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: 
  let 
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
    name = "api-go-templ";
    vendorHash = pkgs.lib.fakeHash;
  in
  {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [ go_1_22 ];
    };
    packages.${system}.default = pkgs.buildGoModule {
      inherit name vendorHash;
      src = ./.;
    };
    app.${system}.default = {
      type = "app";
      description = "Go Api templ";
      start = ''
        ${pkgs.go}/bin/api-go-templ
      '';
    };
  };
}
