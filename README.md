# Wager Wars
## Project Use Cases
### Login
1. User launches the application
1. User is presented with a splash screen while app initializes
1. Once app has been initialized, user is presented with the extended logo (tank + name) for Wager Wars login page with 4 buttons: login with Google, login with Facebook, login with Twitter, and login with Twitch
    1. Clicking on each button will direct the user to login using the corresponding login method for the chosen provider. If the user has linked multiple providers to their account, logging in with any of the linked providers should take them to the same account and initialize a session for all linked providers
1. If login is not successful, repeat previous step. If login is successful, user is taken to a landing page that contains the menus described in the design considerations section below, with the bulk of the view containing an activity feed of pending wagers, pending conflicts, and wager invitations/list of public wagers started by their friends

### Finding Wagers and Friends
1. User inputs text into the search bar in the top menu
1. User is presented with a dropdown list of wagers and (wager war) users that match the text. Each category of user (Twitter, Facebook, etc.) should be clearly grouped together, and the user should have the ability to expand and collapse each category.
    - If the current user is already friended with a search result in the wager wars app, that result should not be displayed in the search results.
    - If the search result is a user, that element in the search result window should include an 'invite' (or 'oppose'?) button. Clicking the invite button should send an opponent request to that user. Clicking elsewhere on that search result element should open the relevant app on the user's device and display the selected user's profile, defaulting to the web browser version instead if the current user doesn't have the relevant app installed
    - If the search result is a wager, the user is taken to a new page that has the details of the wager, including a list of opponents (with the side they're on), the win conditions of the wager, duration of the wager, cost (ante) for the wager, and a button to join either side. Clicking the button should immediately deduct the cost of the wager from the user's war chest and subscribe them to the chosen side. The buttons on the page should then change to text that includes the side they're on, and revisiting this page should have the same result. There should not be a way for users to reverse their decision or back out of a wager.

### Resolving Wagers
1. A user-created wager's duration ends or the user's author visits the wager's page and the user chooses to end the wager
1. If there were no opponents for the wager, the wager's author receives a full refund
1. If there were opponents for the wager, each user subscribed to that wager receives a notification that the wager period has ended, and the wager is placed into a state of conflict
1. While a wager is in conflict, all users subscribed to that wager may place a vote on which side won, and may submit evidence at this time in the form of videos, images, or URLs by visiting the wager's page. All users may switch their vote as many times as they wish during this period
    - If a majority opinion is reached within the 24 hour period, the wager is placed into a pending state
    - If no majority is reached, or fewer than 33% (negotiable, possibly factoring in the weight of both sides) of users subscribed to that wager vote, the wager is immediately placed into a resolved state and nobody receives any winnings
1. While a wager is in a pending state, any user on the losing side may request a veto by visiting the wager's page and clicking a button. This will put the wager into a state of appeal
1. While a wager is in a state of appeal, all users subscribed to that wager have an additional 24 hours to vote on the result again AND invite any of their friends to vote on the result (possibly offering bribes?), again submitting evidence in the form of videos, images, or URLs if need be. Each vote from a user that did not participate in the original wager will equal 3/5ths of a vote from an original participant, up to a yet to be determined maximum. If a majority opinion was reached in favor of the previous losing side after 24 hours, then the pot will be cut in half. Once the 24 hours of appeal ends, the wager will be placed into a resolved state. Users invited to resolve the conflict should receive an award of some sort
1. Each user on the winning side of a resolved wager should receive an equal share of the pot (possibly minus a percentage, but never less than the original ante?)

## Design Considerations
- Primary colors are blood red and gold doubloon gold, with silver ingot as a secondary

- Logo will be a tank shooting gold coin(s)

- Splash screen should include a gif of a tank rolling over terrain, shooting objects that immediately burst into gold and silver coins. The background of the screen should be red, and slowly change (or drip) with gold as progress nears completion

- The top of the app should include a floating menu with a hamburger icon, search bar with text 'Find wagers and opponents', and gold coin icon with text displaying current funds
  - Clicking on the hamburger icon will cause the hamburger to rotate 180 degrees clockwise, hide the top menu, and transform into a larger menu that includes options named:
    1. War Chest
        - Shows current funds and various ways to gain more money
        - Gaining money can be accomplished by:
            1. Creating and subsequently winning wagers
            1. Voting on the results of wagers that are in appeal
            1. Logging into the app every 24 hours to receive a free allotments of funds
            1. Earning achievements such as winning 1, 10, 100, and 1000 wagers, resolving 1, 10, 100, and 1000 conflicts, etc.
    1. Create Wager
        - Allows user to select between public and private, the friends to be invited, the timeline of the wager, and description of the wager and its win conditions
    1. Pending Wagers
    1. Past Wagers
    1. Conflict Resolution
    1. Opponents
        - A list of friends, ranked by war chest with a button to unfriend and the ability to search
    1. Achievements
    1. Account
        - Menu to link Facebook, Twitter, Google, and Twitch accounts. Whichever they initially used to login should already be selected. Remaining should explain the benefits of linking each type of account
        - Option to disable push notifications
    1. Donate
        - Wager Wars will always be completely free and ad free. Donate if you'd like to support the project
        - Allows donation to PayPal using PayPal and various credit card types
        
## Unknowns
- How to prevent users from just creating dummy accounts to farm
- How to reward users for voting on appealed wagers that's fair to everybody and doesn't easily create a way to farm
- Maximum weight for a group of friends from a user for an appealed wager
- Possibility of 'verified' wagers that whose outcomes and prize money are determined by us
