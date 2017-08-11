package main

import "fmt"

func main() {
	var name string
	var feelings string
	var whatis string
	// Greeting
	//------------------------------------------------------
	fmt.Println("Please Enter your name")
	fmt.Scan(&name)
	fmt.Println("Hello", name, "how are you?")
	fmt.Scan(&feelings)
	switch feelings {
	case "great", "good", "Great", "Good", "Awsome", "awsome":
		fmt.Println("Great, lets start")
	case "bad", "awfull", "sucky":
		fmt.Println("well thats too bad")
		fallthrough
	default:
		fmt.Println("your day is about to get even lot better")
	}
	fmt.Println("Alright my turn to guess")
	//Animal route
	//--------------------------------------------------------------
	fmt.Println("Is it a pet?")
	fmt.Scan(&whatis)
	if whatis == "yes" {
		fmt.Println("Sweet!")
		// pets route ----------------------------------------------------
		fmt.Println("Is it a dog?")
		fmt.Scan(&whatis)
		if whatis == "yes" {
			fmt.Println("I win")
		} else {
			fmt.Println("Is it a cat?")
			fmt.Scan(&whatis)
			if whatis == "yes" {
				fmt.Println("I win!")
			} else {
				fmt.Println("Is it a fish?")
				fmt.Scan(&whatis)
				if whatis == "yes" {
					fmt.Println("I Win")
				} else {
					fmt.Println("Is it a lizzard?")
					fmt.Scan(&whatis)
					if whatis == "yes" {
						fmt.Println("I win!")
					} else {
						fmt.Println("Is it a turtle?")
						fmt.Scan(&whatis)
						if whatis == "yes" {
							fmt.Println("I win!")
						} else {
							fmt.Println("Is it a bird?")
							fmt.Scan(&whatis)
							if whatis == "yes" {
								fmt.Println("I win!")
							} else {
								fmt.Println("Is it a rabit?")
								fmt.Scan(&whatis)
								if whatis == "yes" {
									fmt.Println("I win!")
								} else {
									fmt.Println("Is it a hamster?")
									fmt.Scan(&whatis)
									if whatis == "yes" {
										fmt.Println("I win!")
									} else {
										fmt.Println("Is it a guinea pig?")
										fmt.Scan(&whatis)
										if whatis == "yes" {
											fmt.Println("I win!")
										} else {
											fmt.Println("Is it a mouse?")
											fmt.Scan(&whatis)
											if whatis == "yes" {
												fmt.Println("I win!")
											} else {
												fmt.Println("Is it a rat?")
												fmt.Scan(&whatis)
												if whatis == "yes" {
													fmt.Println("I win!")
												} else {
													fmt.Println("Is it a chinchilla?")
													fmt.Scan(&whatis)
													if whatis == "yes" {
														fmt.Println("I win!")
													} else {
														fmt.Println("You win or cheated")
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}

		}
	} else {
		fmt.Println("dang it")
		fmt.Println("is it a farm animal")
		fmt.Scan(&whatis)
		if whatis == "yes" {
			fmt.Println("Sweet!")
			fmt.Println("Is it a cow?")
			fmt.Scan(&whatis)
			if whatis == "yes" {
				fmt.Println("I win")
			} else {
				fmt.Println("Is it a pig?")
				fmt.Scan(&whatis)
				if whatis == "yes" {
					fmt.Println("I win!")
				} else {
					fmt.Println("Is it a horse?")
					fmt.Scan(&whatis)
					if whatis == "yes" {
						fmt.Println("I win!")
					} else {
						fmt.Println("Is it a chicken?")
						fmt.Scan(&whatis)
						if whatis == "yes" {
							fmt.Println("I win!")
						} else {
							fmt.Println("Is it a turkey?")
							fmt.Scan(&whatis)
							if whatis == "yes" {
								fmt.Println("I win!")
							} else {
								fmt.Println("Is it a mule?")
								fmt.Scan(&whatis)
								if whatis == "yes" {
									fmt.Println("I win")
								} else {
									fmt.Println("Is it a rabit?")
									fmt.Scan(&whatis)
									if whatis == "yes" {
										fmt.Println("I win!")
									} else {
										fmt.Println("Is it a lamb?")
										fmt.Scan(&whatis)
										if whatis == "yes" {
											fmt.Println("I win!")
										} else {
											fmt.Println("Is it a duck?")
											fmt.Scan(&whatis)
											if whatis == "yes" {
												fmt.Println("I win!")
											} else {
												fmt.Println("Is it a mouse?")
												fmt.Scan(&whatis)
												if whatis == "yes" {
													fmt.Println("I win")
												} else {
													fmt.Println("Is it a sheep?")
													fmt.Scan(&whatis)
													if whatis == "yes" {
														fmt.Println("I win!")
													} else {
														fmt.Println("Is it a goat?")
														fmt.Scan(&whatis)
														if whatis == "yes" {
															fmt.Println("I win!")
														} else {
															fmt.Println("You win or cheated")
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			fmt.Println("dang it")
			fmt.Println("is it an animal that lives by the ocean")
			fmt.Scan(&whatis)
			if whatis == "yes" {
				fmt.Println("Sweet!")
				fmt.Println("Is it a fish?")
				fmt.Scan(&whatis)
				if whatis == "yes" {
					fmt.Println("I win")
				} else {
					fmt.Println("Is it a shark?")
					fmt.Scan(&whatis)
					if whatis == "yes" {
						fmt.Println("I win!")
					} else {
						fmt.Println("Is it a whale?")
						fmt.Scan(&whatis)
						if whatis == "yes" {
							fmt.Println("I win!")
						} else {
							fmt.Println("Is it a Jellyfish?")
							fmt.Scan(&whatis)
							if whatis == "yes" {
								fmt.Println("I win!")
							} else {
								fmt.Println("Is it a squid?")
								fmt.Scan(&whatis)
								if whatis == "yes" {
									fmt.Println("I win!")
								} else {
									fmt.Println("Is it a octopus?")
									fmt.Scan(&whatis)
									if whatis == "yes" {
										fmt.Println("I win")
									} else {
										fmt.Println("Is it a seal?")
										fmt.Scan(&whatis)
										if whatis == "yes" {
											fmt.Println("I win!")
										} else {
											fmt.Println("Is it a Walrus?")
											fmt.Scan(&whatis)
											if whatis == "yes" {
												fmt.Println("I win!")
											} else {
												fmt.Println("Is it a Narwhal?")
												fmt.Scan(&whatis)
												if whatis == "yes" {
													fmt.Println("I win!")
												} else {
													fmt.Println("Is it a crab?")
													fmt.Scan(&whatis)
													if whatis == "yes" {
														fmt.Println("I win")
													} else {
														fmt.Println("Is it a lobster?")
														fmt.Scan(&whatis)
														if whatis == "yes" {
															fmt.Println("I win!")
														} else {
															fmt.Println("Is it a clam?")
															fmt.Scan(&whatis)
															if whatis == "yes" {
																fmt.Println("I win!")
															} else {
																fmt.Println("Is it a seagul?")
																fmt.Scan(&whatis)
																if whatis == "yes" {
																	fmt.Println("I win")
																} else {
																	fmt.Println("Is it a pelican?")
																	fmt.Scan(&whatis)
																	if whatis == "yes" {
																		fmt.Println("I win!")
																	} else {
																		fmt.Println("Is it a coral?")
																		fmt.Scan(&whatis)
																		if whatis == "yes" {
																			fmt.Println("I win!")
																		} else {
																			fmt.Println("Is it a seaslug?")
																			fmt.Scan(&whatis)
																			if whatis == "yes" {
																				fmt.Println("I win!")
																			} else {
																				fmt.Println("Is it a seasnake?")
																				fmt.Scan(&whatis)
																				if whatis == "yes" {
																					fmt.Println("I win!")
																				} else {
																					fmt.Println("Is it a cuddle fish?")
																					fmt.Scan(&whatis)
																					if whatis == "yes" {
																						fmt.Println("I win")
																					} else {
																						fmt.Println("Is it a peanguin?")
																						fmt.Scan(&whatis)
																						if whatis == "yes" {
																							fmt.Println("I win!")
																						} else {
																							fmt.Println("Is it a sea turtle?")
																							fmt.Scan(&whatis)
																							if whatis == "yes" {
																								fmt.Println("I win!")
																							} else {
																								fmt.Println("Is it a sea otter?")
																								fmt.Scan(&whatis)
																								if whatis == "yes" {
																									fmt.Println("I win!")
																								} else {

																									fmt.Println("You win or cheated")
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				fmt.Println("dang it")
				fmt.Println("is it an animal that lives on the savana")
				fmt.Scan(&whatis)
				if whatis == "yes" {
					fmt.Println("Sweet!")
					fmt.Println("Is it a Lion?")
					fmt.Scan(&whatis)
					if whatis == "yes" {
						fmt.Println("I win")
					} else {
						fmt.Println("Is it a Gazell?")
						fmt.Scan(&whatis)
						if whatis == "yes" {
							fmt.Println("I win!")
						} else {
							fmt.Println("Is it a Crocidile?")
							fmt.Scan(&whatis)
							if whatis == "yes" {
								fmt.Println("I win!")
							} else {
								fmt.Println("Is it a wildabeast?")
								fmt.Scan(&whatis)
								if whatis == "yes" {
									fmt.Println("I win!")
								} else {
									fmt.Println("Is it a elephant?")
									fmt.Scan(&whatis)
									if whatis == "yes" {
										fmt.Println("I win!")
									} else {
										fmt.Println("Is it a Girrafe?")
										fmt.Scan(&whatis)
										if whatis == "yes" {
											fmt.Println("I win")
										} else {
											fmt.Println("Is it a hiyena?")
											fmt.Scan(&whatis)
											if whatis == "yes" {
												fmt.Println("I win!")
											} else {
												fmt.Println("Is it a birds?")
												fmt.Scan(&whatis)
												if whatis == "yes" {
													fmt.Println("I win!")
												} else {
													fmt.Print("congrats you suffered through this entire game and won")

												}
											}
										}
									}
								}
							}
						}
					}
				}

			}
		}
	}
}
