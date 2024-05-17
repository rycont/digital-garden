import { defineConfig } from "astro/config";
import { visit } from "unist-util-visit";

import mdx from "@astrojs/mdx";

export default defineConfig({
  markdown: {
    remarkPlugins: [makeLinkTitle],
  },
  integrations: [mdx()],
  output: "static",
});

function makeLinkTitle() {
  return (tree) => {
    visit(tree, "link", (node) => {
      if (!node.children.length) {
        node.children = [
          {
            type: "text",
            value: node.url,
          },
        ];

        if(!node.url.includes("/")) {
            node.url = "/" + node.url;
        }
      }
    });
  };
}
