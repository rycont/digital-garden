import articles from "./articles.json" with { type: "json" };

for(const article of articles) {
  const id = article._id || article.title
  const content = article.content
  const date = article.publishedAt

  const markdownContent = `---
title: ${article.title}
layout: ../layouts/article.astro
date: ${date}
---
# ${article.title}

${content}
`

    await Deno.writeTextFile(`./${id}.md`, markdownContent)
}
