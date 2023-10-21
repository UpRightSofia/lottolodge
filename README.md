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

This revolutionary system is not just another banking product, but an experience. It's a call to those who wish to see their money grow, while also indulging in the thrills of a lottery game. Save smart, play smart, and win big!

#### **Notes**

- The game is designed to be played weekly, with the draw occurring every Sunday.

- The game is currently configured to favor high risk-high reward scenarios for marketing purposes. However, the game can be easily configured to favor a more balanced distribution that has increased common win rewards. **It is possible to allow the user to configure his personal reward distribution from balanced to highly risky with a slider.**

- The game is configured to match the expected return of a regular savings account. However, the required capital to operate such lottery might incur additional costs, so it might make sense to adjust the prices (or the game) to yield a lower expected return.

- Given the assumption that the saving account would have ~6,000$ (4,960$ median, 40,000$ average in US savings account). A user would have ~300 tickets per week, hense 50% chance of winning a prize. 

# Installation

```
cp .env.example .env                       #sets required env variables
find scripts -type f -exec chmod +x {} \;  #enables the scripts in /scripts to be executed
```

### Build and run docker image

`./scripts/build_and_tag_image.sh`

to build the image

`./scripts/start.sh`

to start all relevant containers

`./scripts/stop.sh`

to stop them