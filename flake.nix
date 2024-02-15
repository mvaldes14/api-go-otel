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
    vendorHash = "sha256-5tKAmV0jGwtlRLRbxz7WKkefrhQe5+ZT+P52EsEV6ig=";
  in
  {
    devShells.${system}.default = pkgs.mkShell {
      inputsFrom = [ self.packages.${system}.default ];
    };
    packages.${system}.default = pkgs.buildGo122Module {
      inherit name vendorHash;
      src = ./.;
    };
  };
}
