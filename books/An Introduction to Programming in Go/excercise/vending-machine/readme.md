Exercise: Vending Machine  
Coin: TEN(10), Five(5), TWO(2), ONE(1)  
          T     F         TW      O  
Item: Soft Drink(18), Canned Coffee(12), Drinking Water(7)  
Coin Return - returns all inserted money  

# Criteria  
    Unlimited items  
    Unlimited change  
    Currently inserted money  

1. Buy SD(soft drink) with exact change  
Insert: T, F, TW, O  
Currently inserted money: 18  
Choose: Select SD  
Return: SD  

2. Start adding change but hit coin return  
Insert: T, T, F  
Currently inserted money: 25  
Choose: Coin Return  
Return: T, T, F  

3. Buy CC(canned coffee) without exact change  
Insert: T, T  
Currently inserted money: 20  
Choose: Select CC  
Return: F, TW, O  