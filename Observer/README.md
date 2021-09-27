Sometimes we have to be informed when something has changed.
* Object's field changes.
* Object does something.
* Some external event occurs.


One solution would be to continuously poll the object, for examle, every 100s or something like that. This is not so usefull because most of the times you wont get anything. Best idea would be to get notifications based on events.

Observer: An object that wishes to be informed about events that happen in the system.

Observable: The entity generating the events.