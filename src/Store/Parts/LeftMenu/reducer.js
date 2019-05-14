const initState = {
    "elements": [
        {
            id: 1,
            avatar: "https://www.shareicon.net/download/2015/09/24/106382_add_512x512.png",
            action: "create",
        },
        {
            id: 2,
            avatar: "https://www.softrew.ru/uploads/posts/2017-05/1494582842_kak-otkryt-fayl-cdw-3.png",
            action: "open file",
            param: "unkn"
        },
        {
            id: 3,
            avatar: "http://clipart-library.com/data_images/537597.png",
            action: "execute"
        }
    ]
}

export default function LMI_leftMenu(state = initState, action){
    const {type, payload} = action;

    switch(type){
        default:
            return state;
    }
}