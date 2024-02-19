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
    vendorHash = null;
  in
  {
    devShells.${system}.default = pkgs.mkShell {
      inputsFrom = [ self.packages.${system}.default ];
      nativeBuildInputs = [ pkgs.air ];
    };
    packages.${system}.default = pkgs.buildGo122Module {
      inherit name vendorHash;
      src = ./.;
    };
    apps.${system}.default = {
      type = "app";
      program = "${self.devShells.${system}.default.nativeBuildInputs}";
    };
  };
}
