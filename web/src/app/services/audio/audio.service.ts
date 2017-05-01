import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Rx';
import { WebsocketService } from '../websocket/websocket.service';

const CHAT_URL = 'ws://localhost:7080';

export interface Message {
  text: string,
  audio: HTMLAudioElement,
}

@Injectable()
export class AudioService {
  public messages: Subject<Message>;

  constructor(wsService: WebsocketService) {
    this.messages = <Subject<Message>>wsService
      .connect(CHAT_URL)
      .map((response: MessageEvent): Message => {
        let data = JSON.parse(response.data);
        return {
          text:   data.text,
          audio:  new Audio("data:audio/x-wav;base64, " + data.base64),
        }
      });
  }
}
