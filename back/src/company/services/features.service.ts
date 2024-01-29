import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, ObjectId } from 'mongoose';
import { KeyValue } from 'src/shared/types/shared-types';
import { Company } from '../schemas/company.schema';

@Injectable()
export class FeaturesService {
  constructor(
    @InjectModel(Company.name) private companyModel: Model<Company>
  ) {}

  async getFeatures(companyId: string): Promise<KeyValue<string>[]> {
    const company = await this.getCompany(companyId);

    return company.featureFlags;
  }

  async addFeature(
    companyId: string,
    featureName: string
  ): Promise<KeyValue<string>> {
    const company = await this.getCompany(companyId);

    const featureFlag = { key: featureName, value: 'on' };
    company.featureFlags.push(featureFlag);
    company.save();
    return featureFlag;
  }

  async deleteFeature(
    companyId: string,
    featureName: string
  ): Promise<KeyValue<string>> {
    const company = await this.getCompany(companyId);

    const featureFlagIndex = company.featureFlags.findIndex(
      flag => flag.key === featureName
    );

    const featureFlag = company.featureFlags[featureFlagIndex];

    if (featureFlagIndex === -1) {
      throw 'Feature not found';
    }

    company.featureFlags.splice(featureFlagIndex, 1);
    company.save();

    return featureFlag;
  }

  private companyNotFoundError(companyId: string): Error {
    return new Error(`Company with id ${companyId} not found`);
  }

  private async getCompany(companyId: string) {
    const company = await this.companyModel.findById(companyId).exec();

    if (!company) {
      throw this.companyNotFoundError(companyId);
    }
    return company;
  }
}
