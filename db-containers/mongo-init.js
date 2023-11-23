db.createUser(
    {
        user: "root",
        pwd: "sofiene123",
        roles: [
            {
                role: "readWrite",
                db: "chatbot-conversations"
            }
        ]
    }
);