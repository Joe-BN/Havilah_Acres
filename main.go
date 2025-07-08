package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/template/html/v2"
  "time"
)

// Post represents a blog post
type Post struct {
  Title   string
  Date    string
  Excerpt string
  Image   string
  Slug    string
}

func main() {
  engine := html.New("./site/pages", ".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })

  // static assets
  app.Static("/css", "./site/assets/css")
  app.Static("/js", "./site/assets/js")
  app.Static("/images", "./site/images")

  // sample posts
  posts := []Post{
    {
      Title:   "5 Tips for Buying Land in Nakuru",
      Date:    time.Now().AddDate(0, 0, -3).Format("Jan 2, 2006"),
      Excerpt: "Learn how to navigate pricing, legal checks, and value growth when purchasing land in Nakuru.",
      Image:   "/images/land-tips.jpg",
      Slug:    "5-tips-for-buying-land-in-nakuru",
    },
  }

  // Blog index
  app.Get("/blog", func(c *fiber.Ctx) error {
    return c.Render("blog", fiber.Map{
      "Posts": posts,
    })
  })

  // Individual post (stub)
  app.Get("/blog/:slug", func(c *fiber.Ctx) error {
    slug := c.Params("slug")
    // find post by slug...
    return c.SendString("Render post: " + slug)
  })

  app.Listen(":3000")
}