import React from 'react';

import './style.scss';

export default function EmptyField(props){
    return (
        <div 
            className="empty_field" 
            style={{
                background: "#d35400", 
                width: '100%', 
                height: '100%', 
                display: 'flex', 
                flexDirection: 'column', 
                justifyContent: 'center', 
                alignItems: 'center',
                color: 'white'}}>
            <h2>Вы пока не открыли ни одного файла</h2>
            <div className="open_modal_window" style={{marginTop: "45px"}}>
                <h4>Вы можете </h4>
                <div className="open_file_modal_window _button_open">Открыть файл</div>
                <div className="create_new_file_modal_window _button_open">Создать файл</div>
            </div>
        </div>
    )
}