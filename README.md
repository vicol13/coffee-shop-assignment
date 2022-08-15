This repository represents the task for Aaha Inc which was done by valeriu vicol. Bellow you can find the description of the tak




### "Very Popular Coffee Shop"

There was a coffee shop that was so popular, so it introduced quotas for buying coffee. In general it means, that you can buy no more than X coffees in last Y hours. Different coffee types (Espresso, Americano, Cappuccino, etc) have different quotas. This limitation was not good news for frequent clients, so shop started to give out custom memberships (with increased quotas).

#### Example of coffee shop quotas:

```
Basic (For Everybody): 
    3 Espresso in last 24 hours
    3 Americano in last 24 hours
    1 Cappuccino in last 24 hours

Membership "Coffeelover": 
    5 Espresso in last 24 hours
    5 Americano in last 24 hours
    5 Cappuccino in last 24 hours

Membership "Espresso Maniac"
    5 Espresso in last 60 minutes
    Cappuccino/Americano same as "Basic"
```

#### How quotas are used:
Imagine you're a basic client. You bought your first Espresso at *08:30:00* and your second Espresso at *22:30:00*. Now, if you want 3rd Espresso, you must wait until *08:30:00* of the next day. It wasn't comfortable for your needs, so you decided to become an "Espresso Maniac" member. Now, if you had your Espressos at *08:11:30*, *08:18:00*, *08:29:00*, *08:59:00* and *09:11:00*, you must wait just 30 seconds for being able to buy your sixth Espresso. 

#### Task:
1. Write a Golang application (Http server), that runs the coffeeshop.
2. Declare your quotas (basic, and list of memberships) as configurations. It can be a yaml file or can be just hardcoded.
3. Application's only endpoint is:<br>
    `Buy a coffee (specifying coffee type)`
    + User's data (id and membership type) is given via headers. (How users are authorized - is not a part of the task)
    + Endpoint returns `200 OK` in case if coffee is given, and `429 Too many requests` in case when user can't get his coffee because of quoata limit exceeded. Also, endpoint returns (in header or body response) how much user must wait.
4. Use ANY kind of database/data structure store for saving current users' state. But consider, that application itself can be running via multiple instances (concurrent-safe solution)