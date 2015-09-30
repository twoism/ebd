d := decider.New()

if d.IsAvailable("enable_beta_feature") {
  // Run feature code
} else {
  // Run legacy path
}
