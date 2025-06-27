package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/nikkbh/social-app/internal/store"
)

var usernames = []string{
	"PixelPioneer23",
	"GoGetterX",
	"ByteRider42",
	"QuantumKoala7",
	"CodeCrush99",
	"DevNova11",
	"BugHunterZZ",
	"StackSurfer88",
	"SyncWarden3",
	"LoopWizard17",
	"DataDrift91",
	"FuncTasticX",
	"EchoFalcon5",
	"CacheCraze4",
	"GlitchGuru27",
	"CompileKing",
	"SliceShark9",
	"StructStorm",
	"NilNinja77",
	"ConcurrencyCat",
}

var titles = []string{
	"Concurrency in Go: A Practical Approach",
	"Why Go is the Future of Backend Development",
	"Exploring Interfaces in Golang",
	"Building REST APIs with Go and Gin",
	"Go Modules: Dependency Management Made Easy",
	"Mastering Error Handling in Go",
	"Writing Idiomatic Go Code: Do's and Don'ts",
	"Understanding Goroutines and Channels",
	"From Zero to Deploy: Go Web App in 30 Minutes",
	"Unit Testing in Go: A Comprehensive Guide",
	"Profiling and Optimizing Go Programs",
	"Go Generics: What You Need to Know",
	"Handling JSON in Go Like a Pro",
	"Creating CLI Tools with Cobra in Go",
	"Working with Databases in Go Using GORM",
	"Using Context Effectively in Go",
	"Go vs Rust: A Developer’s Perspective",
	"Scaling Go Microservices in Production",
	"Middleware in Go: Building Your Own",
	"A Deep Dive into the Go Memory Model",
}

var contents = []string{
	"Learn how to write clean and efficient concurrent programs using goroutines and channels.",
	"Discover why Go is gaining popularity among backend developers and how it compares to other languages.",
	"Understand the power of interfaces in Go and how they enable flexible, testable code.",
	"A step-by-step tutorial on building a RESTful API using the popular Gin web framework.",
	"Get started with Go modules and manage dependencies like a pro.",
	"Master error handling in Go with practical examples and best practices.",
	"Follow Go's idioms and conventions to write clear, maintainable code.",
	"Explore the mechanics of goroutines and how they help you write concurrent Go programs.",
	"Quickly build and deploy a basic Go web app from scratch.",
	"Learn how to write effective unit tests in Go and increase your code quality.",
	"Profile and optimize your Go code using built-in tools like pprof.",
	"An introduction to generics in Go, with use cases and syntax examples.",
	"Handle JSON data in Go efficiently with encoding/json and custom unmarshalling.",
	"Build powerful command-line tools with the Cobra library in Go.",
	"Integrate relational databases into your Go applications using GORM.",
	"Use context in Go to manage timeouts, cancellations, and request-scoped data.",
	"Compare Go and Rust across performance, safety, and developer experience.",
	"Architect and scale Go microservices using real-world production techniques.",
	"Understand how to write and apply custom middleware in Go web servers.",
	"Dive deep into Go's memory model and how it impacts performance and safety.",
}

var comments = []string{
	"Great explanation, cleared up my confusion!",
	"This was super helpful, thanks!",
	"Can you write a follow-up on testing?",
	"I wish I had found this post earlier!",
	"Awesome tips, especially the part about goroutines.",
	"The JSON handling section was gold.",
	"More examples would make this even better.",
	"I love how concise and clear this is.",
	"Is this approach production-safe?",
	"This helped me fix a bug in my code!",
	"Subscribed for more Go content!",
	"Can you explain slices in more depth?",
	"This article saved me hours—thank you!",
	"Looking forward to a post on Go routines.",
	"How does this compare to Rust?",
	"I used this in my side project, worked perfectly!",
	"Could you include more benchmarks?",
	"Fantastic writing style!",
	"Just what I needed for my interview prep.",
	"Please make a video version of this!",
}

var tags = []string{
	"golang",
	"concurrency",
	"webdev",
	"restapi",
	"microservices",
	"goroutines",
	"channels",
	"json",
	"testing",
	"cli",
	"gorm",
	"performance",
	"error-handling",
	"generics",
	"go-modules",
	"middleware",
	"http",
	"web-frameworks",
	"deployment",
	"context",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating users: ", err)
			return
		}
	}

	posts := generatePosts(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating posts: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comments: ", err)
			return
		}
	}
	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: fmt.Sprintf("%s_%d", usernames[rand.IntN(len(usernames))], i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Password: "1234",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.IntN(len(users))]
		numTags := rand.IntN(5) + 1 // 1 to 5 tags
		tagIndices := rand.Perm(len(tags))[:numTags]
		postTags := make([]string, numTags)
		for j, idx := range tagIndices {
			postTags[j] = tags[idx]
		}
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.IntN(len(titles))],
			Content: contents[rand.IntN(len(contents))],
			Tags:    postTags,
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.IntN(len(posts))].ID,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: comments[rand.IntN(len(comments))],
		}
	}
	return cms
}
