import { defineConfig } from "astro/config";
import { visit } from "unist-util-visit";
import mdx from "@astrojs/mdx";
import sitemap from "@astrojs/sitemap";

// https://astro.build/config
export default defineConfig({
  site: "https://garden.postica.app",
  markdown: {
    remarkPlugins: [makeLinkTitle]
  },
  integrations: [mdx(), sitemap()],
  output: "static"
});
function makeLinkTitle() {
  return tree => {
    visit(tree, "link", node => {
      if (!node.children.length) {
        node.children = [{
          type: "text",
          value: node.url
        }];
      }
    });
  };
}