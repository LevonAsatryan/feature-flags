import { Body, Controller, Delete, Get, Param, Post } from '@nestjs/common';
import { FeaturesService } from '../services/features.service';
import { KeyValue } from 'src/shared/types/shared-types';

@Controller('/company/features')
export class FeaturesController {
  constructor(private readonly featuresService: FeaturesService) {}

  @Get('/:companyId')
  getFeatures(
    @Param() param: { companyId: string }
  ): Promise<KeyValue<string>[]> {
    return this.featuresService.getFeatures(param.companyId);
  }

  @Post('/:companyId')
  createFeature(
    @Body() body: { name: string },
    @Param() param: { companyId: string }
  ): Promise<KeyValue<string>> {
    return this.featuresService.addFeature(param.companyId, body.name);
  }

  @Delete('/:companyId')
  deleteFeature(
    @Body() body: { name: string },
    @Param() param: { companyId: string }
  ): Promise<KeyValue<string>> {
    return this.featuresService.deleteFeature(param.companyId, body.name);
  }
}
