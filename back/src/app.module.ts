import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { CompanyModule } from './company/company.module';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://localhost/feature-flags'),
    CompanyModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
