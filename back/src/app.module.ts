import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { CompanyModule } from './company/company.module';
import { FeatureFlagsModule } from './feature-flags/feature-flags.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ContainersModule } from './containers/containers.module';
import typeorm from './config/typeorm';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: [typeorm],
    }),
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (configService: ConfigService) =>
        configService.get('typeorm'),
    }),
    CompanyModule,
    FeatureFlagsModule,
    ContainersModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
