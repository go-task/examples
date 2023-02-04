import $ from "./vendor/jquery.js";

$("#app").html(
  "<h1>Go Web App Example</h1>" +
    "<p>This app was compiled using Task.</p>" +
    "<p>" +
    `It uses <a target="_blank" href="https://esbuild.github.io/">esbuild</a> to bundle JS and CSS assets, ` +
    `and <a target="_blank" href="https://pkg.go.dev/embed">Go embeb</a> to keep ` +
    "everything (assets and templates) inside a single binary." +
    "</p>"
);
