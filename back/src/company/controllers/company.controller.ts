import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { CompanyService } from '../services/company.service';
import { CreateCompanyDTO } from '../DTO/company.dto';
import { Company } from '../schemas/company.schema';
import { KeyValue } from 'src/shared/types/shared-types';

@Controller('/company')
export class CompanyController {
  constructor(private readonly companyService: CompanyService) {}

  @Post('/create')
  createCompany(@Body() createCompanyDTO: CreateCompanyDTO): Promise<Company> {
    return this.companyService.create(createCompanyDTO);
  }

  @Get('/all')
  findAll(): Promise<Company[]> {
    return this.companyService.findAll();
  }

  @Get()
  getHello(): string {
    return 'Hello World!';
  }
}
