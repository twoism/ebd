d := decider.New()

if d.IsAvailableForId("super_secret_feature", user.id) {
  // Run feature code for 50% of users
} else {
  // Run legacy path
}
