import React, {Component} from 'react';
import {Container, Row, Col, Modal, ModalBody, ModalFooter, ModalHeader} from 'reactstrap';
import ContextMenu from '../ContextMenu';

export default class Navbar extends Component{
    constructor(props){
        super(props);

        this.state = {
            "filemenu": [
                {
                    id: 1,
                    name: "Создать",
                    handler : function(event){
                        document.location.replace('/')
                    }
                },

                {
                    id: 2,
                    name: "Открыть",
                    handler : function(event){
                        
                    }
                },
                {
                    id : 3,
                    name: "Сохранить",
                    handler : function(event){
                        console.log("handle save file: ", event)
                    }
                },
                {
                    id: 4,
                    name: "Сохранить как",
                    handler : function(event){
                        console.log("handle save file as: ", event)
                    }
                }
                ,
                {
                    id: 5,
                    name: "Выход",
                    handler : function(event){
                        console.log("handle exit from application: ", event)
                    }
                }
            ],
            "correctmenu": [
                {
                    id: 6,
                    name: "Отменить"
                }
                ,
                {
                    id: 7,
                    name: "Повторить"
                },
                {
                    id: 8,
                    name: "Вырезать"
                },
                {
                    id: 9,
                    name: 'Копировать'
                },
                {
                    id: 10,
                    name: 'Вставить'
                },
                {
                    id: 11, 
                    name: "Удалить"
                },
                {
                    id: 12, 
                    name: "Выделить всё"
                }
            ],
            "helpmenu": [
                {
                    id: 13,
                    name: "О программе"
                },
                {
                    id: 14,
                    name: "Справка"
                }
            ],
            "modalOpenFile": false
        }

        this.toggle = this.toggle.bind(this);
    }

    toggle(){
        this.setState({
            modalOpenFile: !this.state.modalOpenFile
        });
    }

    render = ()=> {
        return (
            <div className="navbar_component" style={{height: this.props.height}}>
                <Container>
                    {window.innerWidth < 900 ? 
                    <Row style={{marginBottom: '50px'}}>
                        <div className="navbar_mobile" style={{display: 'flex', flexDirection: 'row'}}>
                            <div className="logotype" style={{backgroundImage: 'url("https://vacancy.newhr.ru/data/images/golang.png")', backgroundSize: 'cover', backgroundRepeat: 'no-repeat', backgroundPosition: 'center center', width: 75, height: 75, marginRight: 90}}></div>
                            <ContextMenu name="Файл">
                                {this.state.filemenu.map(menuitem => 
                                    <div key={menuitem.id}>{menuitem.name}</div>
                                    )}
                            </ContextMenu>
                            <ContextMenu name="Правка">
                                {this.state.correctmenu.map(menuitem => 
                                    <div key={menuitem.id}>{menuitem.name}</div>
                                    )}
                            </ContextMenu>
                            <ContextMenu name="Справка">
                                    {this.state.helpmenu.map(menuitem => 
                                        <div key={menuitem.id}>{menuitem.name}</div>
                                        )}
                                </ContextMenu>
                        </div>
                    </Row>
                    :
                    <Row>
                        <Col md={{size: 6}}>
                            <h2>WebDev Compiler for language aGo</h2>
                        </Col>
                        <Col md={{offset: 1, size: 5}}>
                            <Container>
                                <Row>
                                    <Col>
                                        <ContextMenu name="Файл">
                                            {this.state.filemenu.map(menuitem => 
                                                <div key={menuitem.id} onClick={menuitem.handler}>{menuitem.name}</div>
                                                )}
                                        </ContextMenu>
                                    </Col>
                                    <Col>
                                        <ContextMenu name="Правка">
                                            {this.state.correctmenu.map(menuitem => 
                                                <div key={menuitem.id} onClick={menuitem.handler}>{menuitem.name}</div>
                                                )}
                                        </ContextMenu>
                                    </Col>
                                    <Col>
                                        <ContextMenu name="Справка">
                                            {this.state.helpmenu.map(menuitem => 
                                                <div key={menuitem.id} onClick={menuitem.handler}>{menuitem.name}</div>
                                                )}
                                        </ContextMenu>
                                    </Col>
                                </Row>
                            </Container>
                        </Col>
                    </Row>
                    }
                </Container>
            </div>
        )
    }
}