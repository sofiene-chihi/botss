db.createUser(
    {
        user: "",
        pwd: "",
        roles: [
            {
                role: "readWrite",
                db: "chatbot-conversations"
            }
        ]
    }
);