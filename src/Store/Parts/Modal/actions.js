import { CHANGE_SOURCE_MW, OPEN_MODAL_WINDOW, CLOSE_MW } from "./constants";

export function SetModalHelp(data){
    return {
        type: CHANGE_SOURCE_MW,
        payload: {
            "head": "HELP",
            "body": `<div class="body_of_help">
            <pre>This is page for instruction of our application
                    1. Typing your source code in input field
                    2. Touch on Execute button
                    3. Check output results in output block
            </pre>
            </div>`
        }
    }
}

export function SetModalTask(){
    return {
        type: CHANGE_SOURCE_MW,
        payload: {
            "head": "Course Task staging",
            "body": `<div class="body_of_task_staging">
            <p>Declaring varible types in the C++ language</p>
            </div>`
        }
    }
}


export function SetModalGrammar(){
    return {
        type: CHANGE_SOURCE_MW,
        payload: {
            "head": "Grammar definition",
            "body": `<div class="body_of_grammar" style="text-align: left;">
            <pre>
            expression -> type pointer identificator ;
            type -> int | float
            pointer -> '*'|'*'pointer|e
            identificator -> ('A'|...|'Z'|'a'|...|'z'|num)+|, identificatior
            num -> ('1'|'2'|...|'9'|'0')+|e
            </pre>
            </div>`
        }
    }
}


export function OpenModalWindow(){
    return {
        type: OPEN_MODAL_WINDOW,
        payload: true
    }
}

export function CloseModalWindow(){
    return {
        type: CLOSE_MW,
        payload: false,
    }
}