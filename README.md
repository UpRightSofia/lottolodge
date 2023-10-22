# Banking Meets Lottery

## Overview

### **The Future of Savings**

In a bid to reinvent traditional banking methods and make savings more attractive to the masses, we present an innovative saving account system that seamlessly marries banking with the thrill of lottery games.

#### **Concept**

The conventional saving account offers a fixed interest rate, say 4%, paid periodically to the account holder. Our avant-garde system pivots from this norm, offering account holders a two-pronged approach to interest accrual:

1. A guaranteed 2% yield on the savings.
2. For every $25 in the savings account, the account holder earns a free ticket for our weekly lottery game.

#### **The Game**

The lottery game is an exciting draw of numbers with a chance to significantly multiply winnings:

1. **Number Draw**: From a pool of numbers 1 to 99, six distinct numbers are drawn without replacement.
2. **Multiplier Draw**: Two additional draws occur for a 10x and a 20x multiplier, each selected from the same pool of 1 to 99.
3. **Winning**: A ticket starts winning when it matches at least 3 out of the 6 drawn numbers. The real thrill starts with the multipliers: if your ticket matches the 10x or/and the 20x multiplier, your prize can soar, potentially multiplying your winnings by 10 or 20, or even both!

#### **Why This System?**

- **Engagement**: This system plays into the human inclination towards games of chance, adding an element of excitement to the otherwise mundane process of saving.
- **Guaranteed Return**: The base return stands firm at 2%. Furthermore, the expected return from winnings over a year is projected to be another 2%, making the total average return comparable to a regular savings account.
- **Flexibility**: Our game system is highly customizable. Depending on the risk appetite, the game can be configured to either favor high risk-high reward scenarios or lean towards a more balanced distribution that rewards common wins more frequently.

#### **Conclusion**

This revolutionary system is more than just a banking offering — it's an experience that combines financial prudence with the excitement of chance. The prospect of significant returns will motivate increased savings, especially among those less traditionally inclined to set money aside. Embrace a unique savings journey that merges prudence with the thrill of life changing money.


### **Game Insights**

- **Weekly Draws**: Our game is structured to hold draws every Sunday, maintaining consistency and anticipation among participants.

- **Configurability**: Currently set in a high-risk, high-reward format for engagement, the game also offers flexibility. It can be adjusted to provide a more even distribution of rewards. Additionally, there's potential for introducing user customization, allowing savers to adjust the reward bias to their liking.

- **Financial Balance**: We aim for the game's returns to mirror those of traditional savings accounts. As with any game of chance, operating costs might arise from maintaining reserves to cover winnings. Most lotteries reduce expected reward to cover those risks.

- **Participation Odds**: Considering an average savings of approximately $6,000 (based on US statistics with a median of $4,960 and an average of $40,000), a median participant would gain around 300 tickets weekly. This provides a significant 50% chance of securing a win each week.

# Project Structure

- **Root**: Houses configuration and setup files like `Dockerfile`, `docker-compose.yaml`, and `README.md`. The main entry point for the application is found in `main.go`.
- **scripts**: This directory holds shell scripts for various operations, such as Docker image building and starting/stopping services.
- **src**: Main source code resides here and is further subdivided into:
  - **handlers**: Here, you'll find the logic responsible for handling HTTP requests. This logic is divided based on its functionality, i.e., pool, tickets, and winnings.
  - **models**: This section defines the data models and the interactions with the database. Distinct directories for each entity (like pool, settings, ticket, user, winning) and utility functions are placed here.
  - **migrations**: Houses SQL migration files essential for evolving the database structure.
  - **utils**: A collection of utility functions useful across various parts of the project, especially during integration tests.

Each sub-directory generally comprises a mix of Go source files (`.go`). Unit tests accompany these files and can be recognized by their `_test.go` suffix. SQL scripts (`.sql`) also appear where relevant.

 - Integration tests in models require docker running


# Installation

```
cp .env.example .env                       #sets required env variables
find scripts -type f -exec chmod +x {} \;  #enables the scripts in /scripts to be executed
```

### Test

`./scripts/test.sh`

### Build and run docker image

`./scripts/build_and_tag_image.sh`

to build the image

`./scripts/start.sh`

to start all relevant containers

`./scripts/stop.sh`

to stop them.