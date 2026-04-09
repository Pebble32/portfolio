package templates

import (
	"encoding/xml"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type RSS struct{
	Channel struct{
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct{
      Title      string `xml:"title"`
      AuthorName string `xml:"author_name"`
      UserRating string `xml:"user_rating"`
      UserReview string `xml:"user_review"`
      Link       string `xml:"link"`
      ImageURL   string `xml:"book_image_url"`
      UserShelves string `xml:"user_shelves"`
	  ReadAt string `xml:"user_read_at"`
}

func GetProjectBySlug(slug string) (Project, bool) {
	for _, p := range projects {
		if p.Slug == slug {
			return p, true
		}
	}
	return Project{}, false
}

func GetAndSortBooks() ([]Book, error){
	res, err := http.Get("https://www.goodreads.com/review/list_rss/164496014")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	var books []Book
	for _, item := range rss.Channel.Items{
		shelf := item.UserShelves
		rating, _ := strconv.Atoi(item.UserRating)
		if shelf == "" && (item.ReadAt != "" || rating != 0){
			shelf = "read"
		}
		books = append(books, Book{
			Title:    item.Title,
			Author:   item.AuthorName,
			Rating:   rating,
			Review:   item.UserReview,
			Link:     item.Link,
			ImageURL: item.ImageURL,
			Shelf:    shelf,
		})
	}

	slices.SortFunc(books, func(a, b Book) int {
		return shelfOrder(a.Shelf) - shelfOrder(b.Shelf)
	})
	return books, nil
}

func shelfOrder(shelf string) int{
	switch shelf {
	case "read":
		return 0
	case "currently-reading":
		return 1
	default:
		return 2
	}
}

var projects = []Project{
	{
		Slug:        "itinerary-scoring-pipeline",
		Title:       "Itinerary Scoring ML Pipeline",
		Icon:        "\u2708\uFE0F",
		Description: "ML pipeline processing 2.6 billion rows daily to score and rank flight itineraries at Dohop.",
		LongDescription: "Architected an itinerary scoring ML pipeline achieving 99.4% recall at a 50% index cut, " +
			"processing 2.6 billion rows daily. Optimized content delivery and booking conversion across airline partners. " +
			"Identified and resolved throughput bottlenecks by dropping down to C for performance-critical itinerary " +
			"creation and loading stages of the pipeline.",
		Tags: []string{"Python", "C", "PyTorch", "Distributed Systems"},
	},
	{
		Slug:        "flight-disruption-model",
		Title:       "Flight Disruption Probability Model",
		Icon:        "\u26A0\uFE0F",
		Description: "Near-instant insurance pricing model for flight disruptions, projected to save ~25K/month.",
		LongDescription: "Built a flight disruption probability model serving near-instant insurance pricing lookups at booking time. " +
			"The model predicts the likelihood of delays and cancellations, enabling dynamic pricing for travel insurance products. " +
			"Projected to save between 20-25K per month for the business.",
		Tags: []string{"Python", "PyTorch", "SageMaker"},
	},
	{
		Slug:        "fish-detection-yolo",
		Title:       "Real-Time Fish Species Detection",
		Icon:        "\U0001F41F",
		Description: "YOLO-based fish species and gender detection system, cutting inference time by 66% across 530K frames.",
		LongDescription: "Built a real-time fish species and gender detection system using YOLO architecture at the " +
			"Marine and Freshwater Institute of Iceland. Applied frame-skipping, image downsampling, and a Bayesian " +
			"frame-difference filter to skip redundant frames, cutting inference time by 66% (down to 3.5s) across " +
			"530,000 frames, saving 3M ISK/year.",
		Tags: []string{"Python", "YOLOv5/v8", "PyTorch", "Computer Vision"},
	},
	{
		Slug:        "anomaly-detection",
		Title:       "Booking Anomaly Detection",
		Icon:        "\U0001F6A8",
		Description: "Automated real-time anomaly detection for booking traffic across partners at Dohop.",
		LongDescription: "Built an automated anomaly detection system for real-time monitoring of booking traffic across " +
			"partners and operational parts of the system at Dohop. Deployed to SageMaker Serverless for " +
			"cost-efficient, always-on monitoring without managing infrastructure.",
		Tags: []string{"Python", "SageMaker", "Anomaly Detection", "Real-Time"},
	},
	{
		Slug:        "efficient-3d-cnn",
		Title:       "Efficient 3D CNN for Action Recognition",
		Icon:        "\U0001F3AC",
		Description: "3D CNN achieving 86.24% on UCF-101 with 44.8% fewer parameters than I3D.",
		LongDescription: "Designed an efficient 3D CNN architecture inspired by ResNet-50 bottleneck designs, combined with " +
			"TrivialAugmentation. The model uses approximately 13.8M parameters — 44.8% fewer than I3D — while " +
			"achieving 86.24% accuracy on UCF-101 (vs I3D's 84.5% RGB-only). Trained from scratch, demonstrating " +
			"that careful architectural choices and simple augmentation strategies can yield competitive performance " +
			"without massive parameter counts.",
		Tags: []string{"Python", "PyTorch", "3D CNNs", "Video Classification"},
	},
	{
		Slug:        "video-model-interpretability",
		Title:       "Interpreting Video Models",
		Icon:        "\U0001F50D",
		Description: "Visualizing attention in Video Transformers and Grad-CAM in 3D CNNs for action recognition.",
		LongDescription: "Explored interpretability of video action recognition models. Visualized internal attention " +
			"mechanisms of TimeSFormer to reveal which spatio-temporal regions influence predictions. Applied Grad-CAM " +
			"to 3D CNN-based models, providing insights into how convolutional architectures perceive video content " +
			"across spatial and temporal dimensions.",
		Tags: []string{"Python", "PyTorch", "TimeSFormer", "Grad-CAM"},
	},
	{
		Slug:        "physics-simulation-ml",
		Title:       "Simulation and Learning in Physical Systems",
		Icon:        "\u2699\uFE0F",
		Description: "Neural networks for approximating physical systems: Lorenz attractors, Lotka-Volterra, and the Ising model.",
		LongDescription: "Applied neural networks to learn and simulate physical systems. Covered sine function approximation " +
			"as a baseline, then moved to chaotic dynamics (Lorenz attractor), predator-prey dynamics (Lotka-Volterra), " +
			"and equilibrium/phase transitions via the Ising model using the Metropolis algorithm. Explored how well " +
			"neural networks can capture the behavior of these systems compared to analytical solutions.",
		Tags: []string{"Python", "PyTorch", "Physics Simulation", "Neural Networks"},
	},
	{
		Slug:        "uni-carpooling-app",
		Title:       "University Carpooling Platform",
		Icon:        "\U0001F697",
		Description: "Mobile carpooling app for University of Iceland students and staff with real-time ride matching.",
		LongDescription: "Designed and developed a carpooling platform for Háskóli Íslands, enabling students and staff " +
			"to create and join campus-bound rides from their smartphones. Features include HÍ email verification, " +
			"real-time ride matching, GPS route integration, offline data capture with SQLite sync, and push " +
			"notifications for ride updates. Built to reduce single-occupant commuting and support the university's " +
			"sustainability goals.",
		Tags: []string{"Mobile", "Real-Time", "GPS", "SQLite"},
	},
}
