import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { CompanyModule } from './company/company.module';
import { Company } from './company/entities/company.entity';
import { FeatureFlagsModule } from './feature-flags/feature-flags.module';
import { FeatureFlag } from './feature-flags/entities/feature-flag.entity';
import { ConfigModule } from '@nestjs/config';
import typeorm from './config/typeorm';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: [typeorm],
    }),
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: 'localhost',
      port: 5432,
      password: 'admin',
      username: 'postgres',
      entities: [Company, FeatureFlag],
      database: 'feature-flags',
      synchronize: true,
      logging: true,
    }),
    CompanyModule,
    FeatureFlagsModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
