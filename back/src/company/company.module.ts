import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { Company, CompanySchema } from './schemas/company.schema';
import { CompanyService } from './services/company.service';
import { CompanyController } from './controllers/company.controller';
import { FeaturesController } from './controllers/features.controller';
import { FeaturesService } from './services/features.service';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Company.name, schema: CompanySchema }]),
  ],
  controllers: [CompanyController, FeaturesController],
  providers: [CompanyService, FeaturesService],
})
export class CompanyModule {}
