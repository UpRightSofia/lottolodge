# Number of lottery simulations
n_simulations <- 1

# Number of overall repetitions of the lottery
n_repeats <- 100000
no_winners_count <- 0
wins_sum<-0

for (r in 1:n_repeats) {
  winnings<-0
  winning_tickets_count <- 0
  
  # The winning numbers
  winning_numbers <- sample(1:99, 8)
  yellow<-c(winning_numbers[1:6])
  red<-winning_numbers[7]
  blue<-winning_numbers[8]
  
  
  for (i in 1:n_simulations) {
    ticket <- sample(1:99, 8)
    yellow_ch<-ticket[1:6]
    red_ch<-ticket[7]
    blue_ch<-ticket[8]
    
    matches_yellow <- sum(yellow_ch %in% yellow)
    matches_red <- sum(red_ch %in% yellow)
    matches_blue <- sum(blue_ch %in% yellow)
    
    winnings<-switch(matches_yellow+1,0,0,0,1.5,50,1500,150000)
    if(matches_red==1){
      winnings<-winnings*10
    }
    if(matches_blue==1){
      winnings<-winnings*20
    }
  }
  if (matches_yellow >= 3) {
    winning_tickets_count <- winning_tickets_count + 1
  }
  # Check if no winners in the simulation
  if (winning_tickets_count == 0) {
    no_winners_count <- no_winners_count + 1
  }
  wins_sum<-wins_sum+winnings
}

# Calculate ratio
print(wins_sum)
ratio <- 1- (no_winners_count / n_repeats)

# Print results
print(paste("Number of simulations with no winners: ", no_winners_count))
print(paste("Chance of being a winner: ", ratio))

