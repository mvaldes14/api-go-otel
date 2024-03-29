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
      vendorHash = "sha256-+18AtUrInojsgf32Ukh/iQct0cCoJHatVGNsMbRKJPQ=";
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        inputsFrom = [ self.packages.${system}.default ];
        nativeBuildInputs = [ pkgs.air pkgs.templ pkgs.sqlite ];
      };
      packages.${system}.default = pkgs.buildGo122Module {
        inherit name vendorHash;
        src = ./.;
      };

      app.${system}.default = {
        type = "app";
        program = "${pkgs.air}/bin/air";
      };
    };
}
