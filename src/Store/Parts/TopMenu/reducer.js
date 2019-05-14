const initState = {
    "elements": [
        {
            id: 1,
            name: "File",
            elementsInline: [
                {
                    id: 10,
                    name: "New File",
                    action:"create",
                },
                {
                    id: 11,
                    name: "New Window",
                    action:"new",
                },
                {
                    id: 12,
                    name: "Open File",
                    action: "open"
                },
                {
                    id: 13,
                    name: "Save File",
                    action:"save",
                },
                {
                    id: 14,
                    name: "Exit",
                    action: "exit"
                }
            ]
        },
        {
            id: 2,
            name: "Task",
            elementsInline: [
                {
                    id: 21,
                    name: "staging",
                    action: "modal",
                    params: "stag"
                },
                {
                    id: 22,
                    name: "grammar",
                    action: "modal",
                    params: "gram"
                },
            ]
        },
        {
            id: 3,
            name: "Tests",
            elementsInline: [
                {
                    id: 31,
                    name: "First example",
                    action: "open",
                    params: "example1"
                },
                {
                    id: 32,
                    name: "Second example",
                    action: "open",
                    params: "example2"
                },
                {
                    id: 33,
                    name: "Third example",
                    action: "open",
                    params: "example3"
                },
                {
                    id: 34,
                    name: "Fourth example",
                    action: "open",
                    params: "example4"
                },
                {
                    id: 35,
                    name: "Fifth example",
                    action: "open",
                    params: "example5"
                }
            ]
        },
        {
            id: 4,
            name: "Help",
            action: "modal",
            params: "help",
        }
    ]
}

export default function TMI_topMenu(state = initState, action){
    const {type, payload} = action;

    switch(type){
        default:
            return state;
    }
}