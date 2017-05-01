import { Component } from '@angular/core';
import { WebsocketService } from './services/websocket/websocket.service';
import { AudioService } from './services/audio/audio.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [ WebsocketService, AudioService ]
})

export class AppComponent {
  public current: string;

  constructor(private chatService: AudioService) {
    chatService.messages.subscribe(msg => {
      this.current = msg.text;
      msg.audio.play();
    });
  }
}
